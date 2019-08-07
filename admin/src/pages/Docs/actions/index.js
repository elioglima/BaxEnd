import { RAPI } from '../../src/request-api'

export const dispRAPI = (uri, params)  => dispatch =>  { return RAPI(uri,params) }

export const getName = (value)  => dispatch =>  { 
    console.log(value, value.indexOf("NMCP"), value.length)

    if (value.indexOf("CompReactID") > -1) { 
        if (value.indexOf("NMCP") > -1) { 
            return value.substring(value.indexOf("NMCP")+4, value.length)
        }
    }
    
    return  value
}