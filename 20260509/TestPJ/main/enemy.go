components {
  id: "enemy"
  component: "/main/enemy.script"
}
embedded_components {
  id: "ghost"
  type: "sprite"
  data: "default_animation: \"idle_d\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 16.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ghost.tilesource\"\n"
  "}\n"
  ""
}
