import API_ROUTE from "./apiRoute"
import axios from 'axios'

class InfoTransactionApi {
    static getInfo() {
        const uri = API_ROUTE + "/info"
        return axios.get(uri)
    }
}

export default InfoTransactionApi