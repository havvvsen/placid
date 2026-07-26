import { Track } from "./track";

export interface RegisterResponse {
  message: string;
}

export interface LoginResponse {
  message: string;
  token: string
}

export interface UserResponse {
  uuid: string,
  email: string,
  isAdmin: boolean,
  isPremium: boolean,
  createdAt: string
}


export interface TracksResponse {
  tracks: Track[]
}

export interface SubscribeNewsletterResponse {
  message: string
}
