components {
  id: "shadow"
  component: "/main/shadow.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 16.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/shadow.tilesource\"\n"
  "}\n"
  ""
  position {
    z: -0.3
  }
}
