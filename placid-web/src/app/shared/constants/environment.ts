let apiVersion = 'v1';
let prefix = `api/${apiVersion}`;

// apiBaseUrl: "http://localhost:3000",
// trackServerBaseUrl: "http://localhost:3001/tracks",
// bannerServerBaseUrl: "http://localhost:3001/tracks/images",
export const Environment = {
  apiBaseUrl: 'http://placidus.duckdns.org',
  trackServerBaseUrl: 'http://placidus.duckdns.org/assets/tracks',
  bannerServerBaseUrl: 'http://placidus.duckdns.org/assets/banners',
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
  },
};
