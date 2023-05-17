import {
  GAME_ROUTE,
  LOBBY_ROUTE,
  LOGIN_ROUTE,
  REGISTRATION_ROUTE,
} from "../../utils/consts"
import GamePage from "../game/GamePage"
import LobbyPage from "../lobby/LobbyPage"
import LoginPage from "../login/LoginPage"
import RegistrationPage from "../registration/RegistrationPage"

export const publicRoutes = [
  { path: LOGIN_ROUTE, component: <LoginPage /> },
  { path: REGISTRATION_ROUTE, component: <RegistrationPage /> },
]

export const restrictedRoutes = [
  { path: LOBBY_ROUTE, component: <LobbyPage /> },
  { path: GAME_ROUTE + "/:id", component: <GamePage /> },
]
