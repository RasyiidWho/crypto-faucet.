import React from 'react';
import { Form, 
         FormGroup, 
         FormText, 
         Button, 
         Input 
       } from 'reactstrap'

import QrReader from 'react-qr-scanner'
import "../assets/css/submit-button-style.css"
import "../assets/css/body-style.css"

import { NULL, 
         COND_1, 
         COND_2, 
         LENGTH_VALID_ADDRESS, 
         TIME_DELAY_SCAN,
         SIZE_FRAME,
         THRESHOLD_NUMBER_OF_TRANSACTION,
         TOTAL_CURRENT_TRANSACTION_LOGS,
         INPUT_CONTENT, 
         NOTIFICATION_USER_ENTER_OTHER_ADDRESS,
         NOTIFICATION_USER_FULL_SLOTS,
         NOTIFICATION_USER_ENTER_VALID_ADDRESS,
         SEND_REQUEST_BUTTON,
         DESCRIPTION,
         CRYPTO_CURRENCY
       } from "./constant";


function checkAddress (param) {
    const prefix = param.slice(0, 3)
    return (prefix === COND_1 || prefix === COND_2) 
           && param.length === LENGTH_VALID_ADDRESS
}
class NewTransaction extends React.Component {
    state = {
      delay: TIME_DELAY_SCAN,
      paymentAddress: NULL,
      disabled: true
    };

    handleInputChange = e => {
      this.setState({
        [e.target.name]: e.target.value
      });
    };

    handleSubmit = e => {
      e.preventDefault();
      if (this.state.paymentAddress.trim()) {
        this.props.onAddTransaction(this.state);
        this.handleReset();
      }
    };

    handleError = () => {
        this.handleReset()
    }

    handleScan = (data) => {
      if (data) {
          this.setState({
              paymentAddress: this.state.paymentAddress 
                              ? this.state.paymentAddress : data.text,
              disabled: this.state.paymentAddress 
                              ? false : this.state.disabled
          })
      }
    }

    handleReset = () => {
      this.setState({
          paymentAddress: NULL,
          disabled: true
      });
    };

    render() {
      const previewStyle = {
          height: SIZE_FRAME,
          width: SIZE_FRAME,
        }
      return (
        <div className="form-input-alignment">
          <Form onSubmit={ 
              this.props.info[this.state.paymentAddress] > THRESHOLD_NUMBER_OF_TRANSACTION 
              || this.props.info[TOTAL_CURRENT_TRANSACTION_LOGS] > 10 
              || !checkAddress(this.state.paymentAddress) 
              ? this.handleError : this.handleSubmit 
          }>
              <FormGroup>
                <div className="form-group container">
                    <Input 
                            className="w-75 container"
                            name="paymentAddress" 
                            onChange={this.handleInputChange} 
                            placeholder= {INPUT_CONTENT}
                            value={ this.state.paymentAddress }
                            disabled={true}
                    />
                    <br/>
                    { 
                      this.props.info[this.state.paymentAddress] > 1 
                      ?  NOTIFICATION_USER_ENTER_OTHER_ADDRESS : NULL
                    }
                    { 
                      checkAddress(this.state.paymentAddress) 
                      && this.props.info[TOTAL_CURRENT_TRANSACTION_LOGS] > 10 
                      ?  NOTIFICATION_USER_FULL_SLOTS : NULL}
                    {
                        !checkAddress(this.state.paymentAddress) 
                        && this.state.paymentAddress.length > 0  
                        ? NOTIFICATION_USER_ENTER_VALID_ADDRESS : NULL
                    }
                    { this.state.disabled &&
                      <QrReader
                          delay={this.state.delay}
                          style={previewStyle}
                          onError={this.handleError}
                          onScan={this.handleScan}
                      />
                    }
                    <div className="container">
                        {this.props.info[this.state.paymentAddress] > THRESHOLD_NUMBER_OF_TRANSACTION 
                        || this.props.info["all"] > TOTAL_CURRENT_TRANSACTION_LOGS 
                        || !checkAddress(this.state.paymentAddress) 
                        ? NULL : <Button type="submit" color="primary">{ SEND_REQUEST_BUTTON }</Button>}
                        <FormText>
                            {DESCRIPTION} {this.props.balance}{CRYPTO_CURRENCY})
                        </FormText>
                    </div>  
                </div>
              </FormGroup>
          </Form>
        </div>
      );
    }
}

export default NewTransaction