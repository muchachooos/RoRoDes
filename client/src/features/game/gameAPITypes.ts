export interface ICell {
    x: number;
    y: number;
    "unit_id": string | null
}

export type IField = ICell[]