let apiVersion = "v1"
let prefix = `api/${apiVersion}`

export const Environment = {
  apiBaseUrl: "http://placidus.duckdns.org:3000",
  trackServerBaseUrl: "http://placidus.duckdns.org:3001/tracks",
  bannerServerBaseUrl: "http://placidus.duckdns.org:3001/tracks/images",
  endpoints: {
    register: `${prefix}/register`,
    login: `${prefix}/login`,
    user: `${prefix}/user`,
    deleteAccount: `${prefix}/delete-account`,
    tracks: `${prefix}/tracks`,
    subscribeNewsletter: `${prefix}/subscribe-newsletter`,
    unSubscribeNewsletter: `${prefix}/unsubscribe-newsletter`,
    adminUploadTrack: `${prefix}/admin/upload-track`,
    adminDeleteTrack: `${prefix}/admin/delete-track`,
  }
}

