import { Navigate, Route, Routes } from "react-router-dom"
import { publicRoutes, restrictedRoutes } from "./routes"
import { LOBBY_ROUTE, LOGIN_ROUTE } from "../../utils/routesConsts"

type Props = {}
const AppRouter = (props: Props) => {
  const isAuth = true
  return (
    <Routes>
      {isAuth
        ? restrictedRoutes.map((r) => {
            const { path, component } = r
            return <Route key={path} path={path} element={component} />
          })
        : publicRoutes.map((r) => {
            const { path, component } = r
            return <Route key={path} path={path} element={component} />
          })}
      <Route
        path="*"
        element={<Navigate to={isAuth ? LOBBY_ROUTE : LOGIN_ROUTE} />}
      />
    </Routes>
  )
}
export default AppRouter
