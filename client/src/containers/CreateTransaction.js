import { connect } from 'react-redux'
import { createTransactionRequest } from '../actions/createNewTransactionAction'
import NewTransaction from '../components/NewTransaction'

const mapDispatchToProps = dispatch => {
  return {
    onAddTransaction: paymentAddress => {
      dispatch(createTransactionRequest(paymentAddress))
    }
  }
}

const mapStateToProps = state => {
    return {
      balance: state.balance, 
      info: state.info
    };
  };

export default connect(
        mapStateToProps,
        mapDispatchToProps
      )(NewTransaction)
