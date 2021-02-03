import API_ROUTE from "./apiRoute"
import axios from 'axios'

class BalanceApi {
    static getBalance() {
        const uri = API_ROUTE + "/balance"
        return axios.get(uri)
    }
}

export default BalanceApi