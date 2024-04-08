export type UsersDataType = {
  data: UserType[]
  size: number
}

export type UserType = {
  id: number
  company_id: number
  first_name: string
  last_name: string
  email: string
}
