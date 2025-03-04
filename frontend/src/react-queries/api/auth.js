import { setUserInfo } from "../../utils/localStorage"
import { unAuthAxios } from "../axios"

const URL = "auth"
export async function login(data)  {

    const res = await unAuthAxios  .post(`${URL}/login`, data)

    const token = JSON.parse(JSON.stringify(res.data)).data
    setUserInfo(token)
    if(token) return true

}