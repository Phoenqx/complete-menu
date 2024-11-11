components {
  id: "BrownBirdController"
  component: "/scripts/BrownBirdController.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemyBirdFly\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 250.0\n"
  "  y: 250.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets.atlas\"\n"
  "}\n"
  ""
  scale {
    x: -0.575312
    y: 0.672633
    z: 1.746667
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"obstacle\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 19.066666\n"
  "  data: 12.704924\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
