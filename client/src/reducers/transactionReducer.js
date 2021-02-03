import { ADD_TRANSACTION, FETCH_TRANSACTION } from '../actions/constants';

export default function transactionReducer(state = [], action) {
  switch (action.type) {
    case ADD_TRANSACTION:
      return [...state, action.payload];
    case FETCH_TRANSACTION:
      return action.transactions;
    default:
      return state;
  }
}