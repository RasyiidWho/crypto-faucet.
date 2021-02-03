import { FETCH_BALANCE } from "../actions/constants";

export default function balanceReducer(state = [], action) {
    switch (action.type) {
      case FETCH_BALANCE:
        return action.balance
      default:
        return state;
    }
  }