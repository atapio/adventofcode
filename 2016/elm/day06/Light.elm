module Light (Model, init, Action(..), update, view) where

-- The light is either on or off
type alias Model = Bool

-- Initially it is off
init: Model
init =
  False

-- Can be turned on and off and toggled
type Action
  = TurnOn
  | TurnOff
  | Toggle

-- updating
update : Action -> Model -> Model
update action state =
  case action of
    TurnOn ->
      True
    TurnOff ->
      False
    Toggle ->
      not state == False -- why do we need the == False

view : Model -> String
view model =
  case model of
    True ->
      "on"
    False ->
      "off"
