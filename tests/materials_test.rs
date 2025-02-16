#[test]
fn test_refract() {
    let uv = Vec3::new(1.0, 1.0, 0.0);
    let n = Vec3::new(-1.0, 0.0, 0.0);
    let etai_over_etat = 1.0;
    let expected = Vec3::new(0.0, 1.0, 0.0);
    let actual = refract(&uv, &n, etai_over_etat);
    assert_eq!(actual, expected);
}

#[test]
fn test_reflectance() {
    let cosine = 0.0;
    let ref_idx = 1.5;
    let expected = 1.0;
    let actual = reflectance(cosine, ref_idx);
    assert_eq!(actual, expected);
}

#[test]
fn test_texture() {
    let _world = Material::Texture(Texture::new(
        Srgb::new(1.0, 1.0, 1.0),
        "data/earth.jpg",
        0.0,
    ));
}

#[test]
fn test_to_json() {
    let m = Metal::new(Srgb::new(0.8, 0.8, 0.8), 2.0);
    let serialized = serde_json::to_string(&m).unwrap();
    assert_eq!(r#"{"albedo":[0.8,0.8,0.8],"fuzz":2.0}"#, serialized,);
}
