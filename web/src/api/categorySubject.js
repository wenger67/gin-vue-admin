import service from '@/utils/request'

export const getSubjectById = (data) => {
  return service({
    url: "/categorySubject/getSubjectById",
    method: 'post',
    data
  })
}

export const getSubjectList = (data) => {
  return service({
    url: "/categorySubject/getSubjectList",
    method: 'post',
    data
  })
}

export const createSubject = (data) => {
  return service({
    url: "/categorySubject/createSubject",
    method: 'post',
    data
  })
}

export const updateSubject = (data) => {
  return service({
    url: "/categorySubject/updateSubject",
    method: 'post',
    data
  })
}

export const deleteSubject = (data) => {
  return service({
    url: "/categorySubject/deleteSubject",
    method: 'post',
    data
  })
}


export const getAllSubjects = (data) => {
  return service({
    url: "/categorySubject/getAllSubjects",
    method: 'post',
    data
  })
}
