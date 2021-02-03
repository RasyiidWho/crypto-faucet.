import { FETCH_INFO_COUNT_OF_TRANSACTION } from "../actions/constants";

export default function infoReducer(state = [], action) {
    switch (action.type) {
      case FETCH_INFO_COUNT_OF_TRANSACTION:
        return action.transactionsInfo
      default:
        return state;
    }
}