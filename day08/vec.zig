const std = @import("std");

fn _dist_squared(left: u64, right: u64) u64 {
    const left_u64: i64 = @intCast(left);
    const right_u64: i64 = @intCast(right);
    const res = std.math.pow(i64, right_u64 - left_u64, 2);
    return @intCast(res);
}
pub const Vec3 = struct {
    x: u64,
    y: u64,
    z: u64,

    pub fn dist_squared(self: Vec3, other: Vec3) u64 {
        return _dist_squared(self.x, other.x) +
            _dist_squared(self.y, other.y) + _dist_squared(self.z, other.z);
    }

    /// Parse strings in the format "x,y,z"
    pub fn parse_string(input: []const u8) !Vec3 {
        var coord_it = std.mem.tokenizeScalar(u8, input, ',');

        const x = try std.fmt.parseUnsigned(u64, coord_it.next().?, 10);
        const y = try std.fmt.parseUnsigned(u64, coord_it.next().?, 10);
        const z = try std.fmt.parseUnsigned(u64, coord_it.next().?, 10);

        return Vec3{ .x = x, .y = y, .z = z };
    }
};
