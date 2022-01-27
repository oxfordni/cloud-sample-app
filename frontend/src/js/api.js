export const goPrefix = '/go'
export const pyPrefix = '/py'
export const apiPrefix = '/api/v1'

const api = {
  backendGoStatus: `${goPrefix}/health`,
  backendPyStatus: `${pyPrefix}/health`,
  getRandomQuote: `${apiPrefix}/quotes/movie-quotes`,
}

export default api
