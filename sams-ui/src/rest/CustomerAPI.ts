import restAPI from '@/rest/restAPI'

export const getAllCustomers = async () => {
  const response = await restAPI.get('/customers')

  return response.data
}