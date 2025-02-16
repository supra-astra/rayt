use palette::Srgb;

use crate::ray::{HitRecord, Ray};

pub trait Scatterable {
    fn scatter(&self, ray: &Ray, hit_record: &HitRecord) -> Option<(Option<Ray>, Srgb)>;
}
// https://docs.rs/serde_with/1.9.4/serde_with/macro.serde_conv.html
serde_with::serde_conv!(
    SrgbAsArray,
    Srgb,
    |srgb: &Srgb| [srgb.red, srgb.green, srgb.blue],
    |value: [f32; 3]| -> Result<_, std::convert::Infallible> {
        Ok(Srgb::new(value[0], value[1], value[2]))
    }
);
serde_with::serde_conv!(
    TexturePixelsAsPath,
    Vec<u8>,
    |_pixels: &Vec<u8>| "/tmp/texture.jpg",
    |value: &str| -> Result<_, std::convert::Infallible> { Ok(load_texture_image(value).0) }
);

//material type can be of different types
#[derive(Debug, Clone, Deserialize, Serialize)]
pub enum Material {
    Lambertian(crate::materials::Material),
    Metal(crate::materials::Material),
    Glass(crate::materials::Material),
    Texture(crate::materials::Material),
    Light(crate::materials::Material),
}

impl Scatterable for Material {
    fn scatter(&self, ray: &Ray, hit_record: &HitRecord) -> Option<(Option<Ray>, Srgb)> {
        match self {
            //differnt scatter properties for differnt materials
            Material::Lambertian(l) => l.scatter(ray, hit_record),
            Material::Metal(m) => m.scatter(ray, hit_record),
            Material::Glass(g) => g.scatter(ray, hit_record),
            Material::Texture(t) => t.scatter(ray, hit_record),
            Material::Light(l) => l.scatter(ray, hit_record),
        }
    }
}
