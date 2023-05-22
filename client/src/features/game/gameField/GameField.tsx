import { FC } from "react"
import { IField } from "../gameAPITypes"

type Props = {
  field: IField
}
const GameField: FC<Props> = ({ field }) => {
  return (
    <div className="flex flex-wrap w-80 h-[200px] mx-auto">
      {field.map((cell) => (
        <div
          key={`${cell.x}${cell.y}`}
          className="w-10 h-10 bg-green-300 border-black border-2" 
        ></div>
      ))}
    </div>
  )
}
export default GameField
