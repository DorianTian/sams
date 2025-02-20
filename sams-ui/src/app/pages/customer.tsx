'use client'

import React, { useCallback, useEffect } from 'react'
import { getAllCustomers } from '@/rest/CustomerAPI'

const Customer = () => {

  const queryCustomers = useCallback(async () => {
    const res = await getAllCustomers()

    console.log('customers: ', res)
  }, [])

  useEffect(() => {
    queryCustomers().then(r => console.log('r: ', r))
  }, [])

  return <div>customer </div>
}

export default Customer