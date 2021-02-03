import API_ROUTE from "./apiRoute"
import axios from 'axios'

class TransactionsApi {
    static getTransactions() {
        const uri = API_ROUTE + "/all"
        return axios.get(uri)
    }
}

export default TransactionsApi