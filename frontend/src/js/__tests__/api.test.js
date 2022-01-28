import api from '../api'

describe('testing the api object', () => {
  it('should have the correct properties', () => {
    expect(api).toHaveProperty('backendGoStatus')
    expect(api).toHaveProperty('backendPyStatus')
    expect(api).toHaveProperty('getRandomQuote')
  })
})
