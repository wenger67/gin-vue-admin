import service from '@/utils/request'

export const getSubjectById = (data) => {
  return service({
    url: "/subject/getSubjectById",
    method: 'post',
    data
  })
}

export const getSubjectList = (data) => {
  return service({
    url: "/subject/getSubjectList",
    method: 'post',
    data
  })
}

export const createSubject = (data) => {
  return service({
    url: "/subject/createSubject",
    method: 'post',
    data
  })
}

export const updateSubject = (data) => {
  return service({
    url: "/subject/updateSubject",
    method: 'post',
    data
  })
}

export const deleteSubject = (data) => {
  return service({
    url: "/subject/deleteSubject",
    method: 'post',
    data
  })
}


export const getAllSubjects = (data) => {
  return service({
    url: "/subject/getAllSubjects",
    method: 'post',
    data
  })
}
