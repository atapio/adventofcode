-- module LightGrid (Model, init, Action, update, view) where
-- module LightGrid (Model, init) where
module LightGrid where


import Dict
import List

import Light

type alias Point =
  { x : Int
  , y : Int
}


type alias Model =
  { size : Point
  , lights : Dict.Dict Int Light.Model
}

init : Int -> Int -> Model
init width height =
  { size = { x = width, y = height }
  , lights = Dict.empty
  }

lightsOn : Model -> Int
lightsOn grid =
  List.length (Dict.keys (Dict.filter (\k v -> v == True) grid.lights))

type Action
  = TurnOn Point Point
  | TurnOff Point Point
  | Toggle Point Point

update : Action -> Model -> Model
update action state =
  case action of
    TurnOn from to ->
      let
        points = List.map (\n -> pointIndex state.size n) (pointsBetween state.size from to)
      in
        {state | lights = List.foldr (\l d -> Dict.update l (updateLight Light.TurnOn) d) state.lights points }
    TurnOff from to ->
      let
        points = List.map (\n -> pointIndex state.size n) (pointsBetween state.size from to)
      in
        {state | lights = List.foldr (\l d -> Dict.update l (updateLight Light.TurnOff) d) state.lights points }
    Toggle from to ->
      let
        points = List.map (\n -> pointIndex state.size n) (pointsBetween state.size from to)
      in
        {state | lights = List.foldr (\l d -> Dict.update l (updateLight Light.Toggle) d) state.lights points }


updateLight : Light.Action -> Maybe Light.Model -> Maybe Light.Model
updateLight action l =
  Just (Light.update action (Maybe.withDefault Light.init l))


pointIndex : Point -> Point -> Int
pointIndex gridSize point =
  gridSize.x * point.y + point.x

pointsBetween : Point -> Point -> Point -> List Point
pointsBetween gridSize from to =
  if to.y == from.y
  then
    List.map (\x -> { from | x = x }) [from.x..to.x]
  else
    List.append (pointsBetween gridSize { from | y = to.y } to) (pointsBetween gridSize from { to | y = to.y - 1 })
