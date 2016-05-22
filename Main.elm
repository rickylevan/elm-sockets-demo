import Html
import Html.App as Html
import Html.Attributes exposing (..)
import Html.Events
import WebSocket

main =
  Html.program
    { init = init
    , view = view
    , update = update
    , subscriptions = subscriptions
    }

echoServer : String
echoServer =
  "ws://localhost:9998"


type alias State =
  { boxText : String
  , socketName : String
  , locked : Bool
  }


init : (State, Cmd Msg)
init = (State "" "" True, webSockCmd)

webSockCmd = WebSocket.send echoServer ""

type Msg
  = BoxInput String
  | Click
  | SockInput String

update : Msg -> State -> (State, Cmd Msg)
update msg state =
  case msg of
    BoxInput newInput ->
      ({state | boxText = newInput, locked = False}, 
      if state.locked then webSockCmd else Cmd.none)
    Click -> (state, webSockCmd)
    SockInput str -> ({state | socketName = str}, Cmd.none)

subscriptions : State -> Sub Msg
subscriptions model =
  WebSocket.listen echoServer SockInput

view : State -> Html.Html Msg
view state =
  Html.div []
    [ Html.input [myStyle green lightGreen, Html.Events.onInput BoxInput] []
    , Html.button [buttonStyle state.locked, Html.Events.onClick Click]
        [Html.text (if state.locked then initMsg else regMsg)]
    , Html.div [myStyle blue lightBlue] 
        [Html.text (if state.locked then "" else (newSaying state))]
    ]

initMsg = "Say something above!"
regMsg = "Next!"

newSaying state = state.boxText ++ ", " ++ state.socketName ++ "!"

green = "#007700"
lightGreen = "#d0ffd0"
blue = "#000077"
lightBlue = "#d0d0ff"
lightPurple = "#cc99d0"
lightGray = "#aaaaaa"


buttonStyle locked = 
   style
    [ ("width", "320px")
    , ("padding", "20px")
    , ("font-family", "Arial, Helvetica, sans-serif")
    , ("font-size", "1.5em")
    , ("text-align" , "center")
    , ("background", if locked then lightGray else lightPurple)
    , ("cursor", "pointer")
    ]

myStyle color bgcolor =
  style
    [ ("width", "100%")
    , ("height", "40px")
    , ("padding", "10px 0")
    , ("font-size", "2em")
    , ("text-align", "center")
    , ("color", color)
    , ("background-color", bgcolor)
    , ("font-family", "Arial, Helvetica, sans-serif")
    ]
