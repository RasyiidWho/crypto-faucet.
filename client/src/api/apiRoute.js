let API_ROUTE 
process.env.NODE_ENV === 'development' ? 
    API_ROUTE = 'http://localhost/v1':
    API_ROUTE = '/'
export default API_ROUTE