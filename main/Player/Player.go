components {
  id: "Player"
  component: "/main/Player/Player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"icon_180\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
