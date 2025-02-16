use ray::Ray;
use vec3::{dot, unit_vector, Vec3};

mod ray;
mod vec3;

//hit the sphere to get a basic camera running
fn hit_sphere(center: Vec3, radius: f64, r: &Ray) -> bool {
    let oc = r.origin() - center;
    let a = dot(&r.direction(), &r.direction());
    let b = 2.0 * dot(&oc, &r.direction());
    let c = dot(&oc, &oc) - radius * radius;
    let discriminant = b * b - (4.0 * a * c);
    discriminant > 0.0
}

// Fixed color function
fn color(r: &Ray) -> Vec3 {
    if hit_sphere(Vec3::new(0.0, 0.0, -1.0), 0.5, r) {
        return Vec3::new(1.0, 0.0, 0.0);
    }
    let unit_direction = unit_vector(&r.direction());
    let t = 0.5 * (unit_direction.y() + 1.0);
    (1.0 - t) * Vec3::new(1.0, 1.0, 1.0) + t * Vec3::new(0.5, 0.7, 1.0)
}

fn main() {
    let nx: i64 = 200;
    let ny: i64 = 100;

    print!("P3\n{} {}\n255\n", nx, ny);

    let lower_left_corner = Vec3::new(-2.0, -1.0, -1.0);
    let horizontal = Vec3::new(4.0, 0.0, 0.0);
    let vertical = Vec3::new(0.0, 2.0, 0.0);
    let origin = Vec3::new(0.0, 0.0, 0.0);

    for j in (0..ny).rev() {
        for i in 0..nx {
            // let v = Vec3::new(i as f64 / nx as f64, j as f64 / ny as f64, 0.2);
            // let ir = (255.99 * v.x()) as i64;
            // let ig = (255.99 * v.y()) as i64;
            // let ib = (255.99 * v.z()) as i64;

            let u = i as f64 / nx as f64;
            let v = j as f64 / ny as f64;
            let r = Ray::new(origin, lower_left_corner + u * horizontal + v * vertical);
            let col = color(&r);

            let ir = (255.99 * col.x()) as i64;
            let ig = (255.99 * col.y()) as i64;
            let ib = (255.99 * col.z()) as i64;

            print!("{} {} {}\n", ir, ig, ib);
        }
    }
}
