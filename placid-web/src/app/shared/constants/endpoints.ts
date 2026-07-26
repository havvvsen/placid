class Endpoints {
  base_url: string = 'http://localhost:3000/api/v1';
  register: string = `${this.base_url}/register`;
  login: string = `${this.base_url}/login`;
  user: string = `${this.base_url}/user`;
  deleteAccount: string = `${this.base_url}/delete-account`;
  tracks: string = `${this.base_url}/tracks`;
  joinNewsLetter: string = `${this.base_url}/join-newsletter`;
  adminUploadTracks: string = `${this.base_url}/admin/upload-tracks`;
  adminDeleteTracks: string = `${this.base_url}/admin/delete-tracks`;
}

export default Endpoints;
