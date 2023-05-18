import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { lobbyAPI } from "./lobbyAPI"
import { GAME_ROUTE } from "../../utils/routesConsts"

const avaliableUsers = [
  { userName: "TestUsername1" },
  { userName: "TestUsername2" },
] // TODO: remove after api endpoint is provided

type Props = {}
const LobbyPage = (props: Props) => {
  const { data } = lobbyAPI.useGetAvaliableUsersQuery(
    {},
    { pollingInterval: 1000 },
  )

  const navigate = useNavigate()

  const handleSelectGame = (id: string) => {
    navigate(`${GAME_ROUTE}/${id}`)
  }

  const handleCreateGame = () => {
    
  }

  return (
    <div className="w-screen h-screen bg-slate-200 flex flex-col justify-center items-center p-5">
      <button
        className="bg-sky-200 p-2 text-sky-800 rounded-md mb-5 w-1/2 mx-auto"
        onClick={handleCreateGame}
      >
        Create New Game
      </button>
      <div className=" w-full max-w-lg max-h-[90%] divide-y-4 divide-white rounded-lg overflow-y-scroll">
        {avaliableUsers?.map((user) => (
          <div
            key={user.userName}
            className="bg-slate-500 text-white p-2 w-full text-center text-ellipsis overflow-hidden cursor-pointer"
            onClick={() => handleSelectGame(user.userName)}
          >
            {user.userName}
          </div>
        ))}
      </div>
    </div>
  )
}
export default LobbyPage
