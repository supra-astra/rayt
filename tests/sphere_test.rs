#[test]
fn test_sphere_hit() {
    let center = Point3D::new(0.0, 0.0, 0.0);
    let sphere = Sphere::new(center, 1.0, Material::Glass(Glass::new(1.5)));
    let ray = Ray::new(Point3D::new(0.0, 0.0, -5.0), Point3D::new(0.0, 0.0, 1.0));
    let hit = sphere.hit(&ray, 0.0, f64::INFINITY);
    assert_eq!(hit.unwrap().t, 4.0);
}

fn test_to_json() {
    let sphere = Sphere::new(
        Point3D::new(0.0, 0.0, 0.0),
        1.0,
        Material::Lambertian(Lambertian::new(Srgb::new(
            0.5 as f32, 0.5 as f32, 0.5 as f32,
        ))),
    );
    let serialized = serde_json::to_string(&sphere).unwrap();
    assert_eq!(
        "{\"center\":{\"x\":0.0,\"y\":0.0,\"z\":0.0},\"radius\":1.0,\"material\":{\"Lambertian\":{\"albedo\":[0.5,0.5,0.5]}}}",
        serialized,
    );
    let s = serde_json::from_str::<Sphere>(&serialized).unwrap();
    assert_eq!(sphere.center, s.center);
    assert_eq!(sphere.radius, s.radius);

    let textured_sphere = Sphere::new(
        Point3D::new(0.0, 0.0, 0.0),
        1.0,
        Material::Texture(Texture::new(
            Srgb::new(0.5 as f32, 0.5 as f32, 0.5 as f32),
            "data/earth.jpg",
            0.0,
        )),
    );

    let tserialized = serde_json::to_string(&textured_sphere).unwrap();
    assert_eq!(
        "{\"center\":{\"x\":0.0,\"y\":0.0,\"z\":0.0},\"radius\":1.0,\"material\":{\"Texture\":{\"albedo\":[0.5,0.5,0.5],\"pixels\":\"/tmp/texture.jpg\",\"width\":2048,\"height\":1024,\"h_offset\":0.0}}}",
        tserialized,
    );

    let tex = Texture::new(
        Srgb::new(0.5 as f32, 0.5 as f32, 0.5 as f32),
        "data/earth.jpg",
        0.0,
    );
    let tloadable = "{\"center\":{\"x\":0.0,\"y\":0.0,\"z\":0.0},\"radius\":1.0,\"material\":{\"Texture\":{\"albedo\":[0.5,0.5,0.5],\"pixels\":\"data/earth.jpg\",\"width\":2048,\"height\":1024,\"h_offset\":0.0}}}";
    let loaded = serde_json::from_str::<Sphere>(&tloadable).unwrap();
    match loaded.material {
        Material::Texture(ref t) => {
            assert_eq!(t.pixels, tex.pixels);
        }
        _ => panic!("Wrong material type"),
    }
}
