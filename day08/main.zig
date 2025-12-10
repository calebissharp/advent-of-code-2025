const std = @import("std");
const vec = @import("./vec.zig");
const actual_input = @embedFile("./input.txt");
const test_input =
    \\162,817,812
    \\57,618,57
    \\906,360,560
    \\592,479,940
    \\352,342,300
    \\466,668,158
    \\542,29,236
    \\431,825,988
    \\739,650,466
    \\52,470,668
    \\216,146,977
    \\819,987,18
    \\117,168,530
    \\805,96,715
    \\346,949,466
    \\970,615,88
    \\941,993,340
    \\862,61,35
    \\984,92,344
    \\425,690,689
;

const JunctionBoxSet = std.AutoHashMap(*JunctionBox, void);

const JunctionGraph = struct {
    junction_boxes: []JunctionBox,
    circuits: std.ArrayList(JunctionBoxSet),

    pub fn deinit(self: *JunctionGraph, allocator: std.mem.Allocator) void {
        self.circuits.deinit(allocator);
        allocator.free(self.junction_boxes);
        self.* = undefined;
    }

    pub fn init(allocator: std.mem.Allocator, junction_boxes: []JunctionBox) !JunctionGraph {
        return JunctionGraph{ .junction_boxes = junction_boxes, .circuits = try std.ArrayList(JunctionBoxSet).initCapacity(allocator, junction_boxes.len) };
    }

    pub fn closestPairs(
        self: JunctionGraph,
        allocator: std.mem.Allocator,
        n: usize,
    ) ![]CoordPair {
        if (self.junction_boxes.len < 2) return &[_]CoordPair{};

        // Number of pairs: len * (len - 1) / 2
        const num_pairs = (self.junction_boxes.len * (self.junction_boxes.len - 1)) / 2;
        const result_len = @min(n, num_pairs);

        var pairs = try std.ArrayList(CoordPair).initCapacity(allocator, num_pairs);
        defer pairs.deinit(allocator);

        for (self.junction_boxes, 0..) |*a, i| {
            for (self.junction_boxes[i + 1 ..]) |*b| {
                try pairs.append(allocator, .{
                    .a = a,
                    .b = b,
                    .distance = a.coord.dist_squared(b.coord),
                });
            }
        }

        std.mem.sort(CoordPair, pairs.items, {}, CoordPair.compare);

        const slice = try pairs.toOwnedSlice(allocator);
        return slice[0..result_len];
    }

    fn find_circuit(self: *JunctionGraph, box: *JunctionBox) ?struct { *JunctionBoxSet, usize } {
        for (self.circuits.items, 0..) |*circuit, i| {
            if (circuit.get(box)) |_| {
                return .{ circuit, i };
            }
        }

        return null;
    }

    /// Create circuits by iterating through `pairs`, up to `max_pairs`
    pub fn createCircuits(self: *JunctionGraph, allocator: std.mem.Allocator, pairs: []CoordPair, max_pairs: usize) !?struct { *JunctionBox, *JunctionBox } {
        for (pairs[0..@min(max_pairs, pairs.len)]) |pair| {
            const maybeCircuitLeft = self.find_circuit(pair.a);
            const maybeCircuitRight = self.find_circuit(pair.b);

            // Check if this is the last connection before there's only one circuit left.
            if (self.circuits.items.len <= 2) {
                // a connection within a circuit won't do anything
                if (maybeCircuitLeft != null and maybeCircuitRight != null) {
                    if (maybeCircuitLeft.?.@"0" == maybeCircuitRight.?.@"0") {
                        continue;
                    }
                }

                const lens = try self.circuitSizes(allocator);
                if (lens.len == 2) {
                    if (lens[0] + lens[1] == self.junction_boxes.len) {
                        return .{ pair.a, pair.b };
                    }
                } else if (lens.len == 1) {
                    if (lens[0] + 1 == self.junction_boxes.len) {
                        return .{ pair.a, pair.b };
                    }
                }
            }

            // Box boxes are in a circuit already
            if (maybeCircuitLeft != null and maybeCircuitRight != null) {
                const circuit_left, const left_i = maybeCircuitLeft.?;
                _ = left_i;
                const circuit_right, const right_i = maybeCircuitRight.?;
                // In the same circuit already, so nothing to do
                if (circuit_left == circuit_right) {
                    continue;
                }

                // Each box is in a separate circuit; merge them
                var it = circuit_right.keyIterator();
                while (it.next()) |other| {
                    try circuit_left.put(other.*, {});
                }
                self.circuits.items[right_i].deinit();
                _ = self.circuits.swapRemove(right_i);
            } else if (maybeCircuitLeft) |left| {
                // Only the left box is in a circuit, put the right one in
                const circuit_left, const left_i = left;
                _ = left_i;
                try circuit_left.put(pair.b, {});
            } else if (maybeCircuitRight) |right| {
                // Only the right box is in a circuit, put the left one in
                const circuit_right, const right_i = right;
                _ = right_i;
                try circuit_right.put(pair.a, {});
            } else {
                // Boxes aren't in any circuit yet; add a new one
                var circuit = JunctionBoxSet.init(allocator);
                try circuit.put(pair.a, {});
                try circuit.put(pair.b, {});
                try self.circuits.append(allocator, circuit);
            }
        }

        return null;
    }

    pub fn circuitSizes(self: *JunctionGraph, allocator: std.mem.Allocator) ![]usize {
        const lengths = try allocator.alloc(u64, self.circuits.items.len);
        for (self.circuits.items, lengths) |circuit, *length| {
            length.* = circuit.count();
        }

        std.sort.block(u64, lengths, {}, std.sort.desc(u64));

        return lengths;
    }
};

const JunctionBox = struct {
    coord: vec.Vec3,

    pub fn init(coord: vec.Vec3) JunctionBox {
        return JunctionBox{ .coord = coord };
    }
};

fn collect_boxes(allocator: std.mem.Allocator, buf: []const u8) !JunctionGraph {
    var it = std.mem.tokenizeScalar(u8, buf, '\n');

    var numRows: usize = 0;
    while (it.next()) |_| {
        numRows += 1;
    }
    it.reset();

    var boxes = try std.ArrayList(JunctionBox).initCapacity(allocator, numRows);
    defer boxes.deinit(allocator);

    while (it.next()) |row| {
        try boxes.append(allocator, JunctionBox.init(try vec.Vec3.parse_string(row)));
    }

    return try JunctionGraph.init(allocator, try boxes.toOwnedSlice(allocator));
}

const CoordPair = struct {
    a: *JunctionBox,
    b: *JunctionBox,
    distance: u64,

    pub fn compare(_: void, self: CoordPair, other: CoordPair) bool {
        return self.distance < other.distance;
    }
};

pub fn part1(input: []const u8, max_pairs: usize) !usize {
    const allocator = std.heap.page_allocator;

    var graph = try collect_boxes(allocator, input);
    defer graph.deinit(allocator);

    const pairs = try graph.closestPairs(allocator, max_pairs);
    defer allocator.free(pairs);

    _ = try graph.createCircuits(allocator, pairs, max_pairs);

    const lengths = try graph.circuitSizes(allocator);

    return lengths[0] * lengths[1] * lengths[2];
}

pub fn part2(input: []const u8) !usize {
    const max_pairs = 100_000_000;
    const allocator = std.heap.page_allocator;

    var graph = try collect_boxes(allocator, input);
    defer graph.deinit(allocator);

    const pairs = try graph.closestPairs(allocator, max_pairs);
    defer allocator.free(pairs);

    const last_pair = try graph.createCircuits(allocator, pairs, max_pairs);

    return last_pair.?.@"0".coord.x * last_pair.?.@"1".coord.x;
}

pub fn main() !void {
    const p1 = try part1(actual_input, 1000);
    const p2 = try part2(actual_input);
    std.debug.print("part 1: {}\n", .{p1});
    std.debug.print("part 2: {}\n", .{p2});
}

test "part 1 works" {
    try std.testing.expectEqual(40, try part1(test_input, 10));
    try std.testing.expectEqual(66640, try part1(actual_input, 1000));
}

test "part 2 works" {
    try std.testing.expectEqual(25272, try part2(test_input));
    try std.testing.expectEqual(78894156, try part2(actual_input));
}
