import API_ROUTE from "./apiRoute"
import axios from 'axios'

class CreateTransactionsApi {
    static createTransactions(paymentAddress) {
        const uri = API_ROUTE + `/donate/${paymentAddress}`
        return axios.get(uri)
    }
}

export default CreateTransactionsApi