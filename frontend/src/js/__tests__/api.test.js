import api from '../api'
import randomQuotes from '../random-quote'
import {useApi} from '../useAPI'

describe('testing the api object', () => {
  it('should have the correct properties', () => {
    expect(api).toHaveProperty('backendGoStatus')
    expect(api).toHaveProperty('backendPyStatus')
    expect(api).toHaveProperty('getRandomQuote')
  })
})

describe('testing the random quotes', () => {
  it('should have the correct properties', () => {
    expect(randomQuotes).toHaveProperty('whyDidYouRender')
  })
})
