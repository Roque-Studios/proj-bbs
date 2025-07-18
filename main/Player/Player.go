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
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"walls\"\n"
  "mask: \"ball\"\n"
  "mask: \"IA\"\n"
  "mask: \"Piss\"\n"
  "mask: \"Poop\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 66.037735\n"
  "  data: 63.181816\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
