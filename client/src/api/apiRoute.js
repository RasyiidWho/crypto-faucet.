let API_ROUTE 
process.env.NODE_ENV === 'development' ? 
    API_ROUTE = 'http://localhost:8085/v1':
    API_ROUTE = 'http://faucet/v1'
export default API_ROUTE