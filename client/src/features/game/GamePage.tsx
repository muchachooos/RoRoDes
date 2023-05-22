import { useEffect } from "react"
import { useParams } from "react-router-dom"
import { gameAPI } from "./gameAPI"
import GameField from "./gameField/GameField"

type Props = {}
const GamePage = (props: Props) => {
  const gameId = useParams().id

  const [getField, { isError, isLoading, data: field }] =
    gameAPI.useLazyGetGameFieldsQuery()

  //   useEffect(() => {
  //     if (gameId){
  //       getField(gameId)
  //     }
  //   }, [gameId])

  if (isLoading) {
    ;<h1>Loading...</h1>
  }
  return (
    <div className="flex justify-center items-center w-screen h-screen bg-slate-300">
      {mockGameFields && <GameField field={mockGameFields} />}
    </div>
  )
}
export default GamePage

//replace 'mockGameFields' with 'field'
var mockGameFields = [
  {
    y: 0,
    x: 0,
    unit_id: null,
  },
  {
    y: 1,
    x: 0,
    unit_id: null,
  },
  {
    y: 2,
    x: 0,
    unit_id: null,
  },
  {
    y: 3,
    x: 0,
    unit_id: null,
  },
  {
    y: 4,
    x: 0,
    unit_id: null,
  },
  {
    y: 0,
    x: 1,
    unit_id: null,
  },
  {
    y: 1,
    x: 1,
    unit_id: null,
  },
  {
    y: 2,
    x: 1,
    unit_id: null,
  },
  {
    y: 3,
    x: 1,
    unit_id: null,
  },
  {
    y: 4,
    x: 1,
    unit_id: null,
  },
  {
    y: 0,
    x: 2,
    unit_id: null,
  },
  {
    y: 1,
    x: 2,
    unit_id: null,
  },
  {
    y: 2,
    x: 2,
    unit_id: null,
  },
  {
    y: 3,
    x: 2,
    unit_id: null,
  },
  {
    y: 4,
    x: 2,
    unit_id: null,
  },
  {
    y: 0,
    x: 3,
    unit_id: null,
  },
  {
    y: 1,
    x: 3,
    unit_id: null,
  },
  {
    y: 2,
    x: 3,
    unit_id: null,
  },
  {
    y: 3,
    x: 3,
    unit_id: null,
  },
  {
    y: 4,
    x: 3,
    unit_id: null,
  },
  {
    y: 0,
    x: 4,
    unit_id: null,
  },
  {
    y: 1,
    x: 4,
    unit_id: null,
  },
  {
    y: 2,
    x: 4,
    unit_id: null,
  },
  {
    y: 3,
    x: 4,
    unit_id: null,
  },
  {
    y: 4,
    x: 4,
    unit_id: null,
  },
  {
    y: 0,
    x: 5,
    unit_id: null,
  },
  {
    y: 1,
    x: 5,
    unit_id: null,
  },
  {
    y: 2,
    x: 5,
    unit_id: null,
  },
  {
    y: 3,
    x: 5,
    unit_id: null,
  },
  {
    y: 4,
    x: 5,
    unit_id: null,
  },
  {
    y: 0,
    x: 6,
    unit_id: null,
  },
  {
    y: 1,
    x: 6,
    unit_id: null,
  },
  {
    y: 2,
    x: 6,
    unit_id: null,
  },
  {
    y: 3,
    x: 6,
    unit_id: null,
  },
  {
    y: 4,
    x: 6,
    unit_id: null,
  },
  {
    y: 0,
    x: 7,
    unit_id: null,
  },
  {
    y: 1,
    x: 7,
    unit_id: null,
  },
  {
    y: 2,
    x: 7,
    unit_id: null,
  },
  {
    y: 3,
    x: 7,
    unit_id: null,
  },
  {
    y: 4,
    x: 7,
    unit_id: null,
  },
]
