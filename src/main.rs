fn main() {
    let nx: i64 = 200;
    let ny: i64 = 100;

    print!("P3\n{} {}\n255\n", nx, ny);

    for j in (0..ny).rev() {
        for i in 0..nx {
            let r = i as f64 / nx as f64;
            let g = j as f64 / ny as f64;
            let b = 0.2;

            let ir = (255.99 * r) as i64;
            let ig = (255.99 * g) as i64;
            let ib = (255.99 * b) as i64;

            print!("{} {} {}\n", ir, ig, ib);
        }
    }
}
