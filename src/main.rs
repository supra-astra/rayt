use vec3::Vec3;

mod vec3;

fn main() {
    let nx: i64 = 200;
    let ny: i64 = 100;

    print!("P3\n{} {}\n255\n", nx, ny);

    for j in (0..ny).rev() {
        for i in 0..nx {
            let v = Vec3::new(i as f64 / nx as f64, j as f64 / ny as f64, 0.2);
            let ir = (255.99 * v.x()) as i64;
            let ig = (255.99 * v.y()) as i64;
            let ib = (255.99 * v.z()) as i64;

            print!("{} {} {}\n", ir, ig, ib);
        }
    }
}
