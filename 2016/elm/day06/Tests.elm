-- Tests.elm
import String
import Graphics.Element exposing (Element)

import ElmTest exposing (..)

import Light
import LightGrid

tests : Test
tests =
  suite "All"
    [ lightSuite
    , lightGridSuite
    ]

lightSuite =
  let
    offLight = Light.update Light.TurnOn Light.init
    onLight = Light.init
  in
    suite "Light"
      [ test "Light is initialized off" (assertEqual False Light.init)
      , test "Light can be turned on" (assertEqual True (Light.update Light.TurnOn offLight))
      , test "Light can be turned off" (assertEqual False (Light.update Light.TurnOff onLight))
      , test "Light can be toggled" (assertEqual False (Light.update Light.Toggle onLight))
      , test "Light can be toggled" (assertEqual True (Light.update Light.Toggle offLight))
      ]

lightGridSuite =
  suite "LightGrid"
  [ test "Point Index is calculated correctly" (assertEqual 3 (LightGrid.pointIndex {  x = 4, y = 4 } { x = 3, y = 0 } ))
  , test "Point Index is calculated correctly" (assertEqual 5050 (LightGrid.pointIndex {  x = 100, y = 100 } { x = 50, y = 50 } ))
  , test "Points between is calculated correctly 1" (assertEqual 1000000 (List.length (LightGrid.pointsBetween {x = 1000, y = 1000} {x=0, y=0} {x=999,y=999})))
  , test "Points between is calculated correctly 2" (assertEqual 1000 (List.length (LightGrid.pointsBetween {x = 1000, y = 1000} {x=0, y=0} {x=999,y=0})))
  , test "Points between is calculated correctly 3" (assertEqual 4 (List.length (LightGrid.pointsBetween {x = 1000, y = 1000} {x=499, y=499} {x=500,y=500})))
  , test "Lights On is calculated correctly" (assertEqual 0 (LightGrid.lightsOn (LightGrid.init 100 100)))
  , test "LightGrid updates properly" (assertEqual 4 (LightGrid.lightsOn (LightGrid.update (LightGrid.TurnOn {x=499, y=499} {x=500,y=500}) (LightGrid.init 100 100))))
  , test "LightGrid updates properly" (assertEqual 1000 (LightGrid.lightsOn (LightGrid.update (LightGrid.TurnOn {x=0, y=0} {x=999,y=0}) (LightGrid.init 100 100))))
  ]


main : Element
main =
    elementRunner tests
