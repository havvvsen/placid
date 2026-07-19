import { Mood } from './mood';

export default class SoundScape {
  bgUrl: string;
  mood: Mood;
  audioUrl: string;

  constructor(bgUrl: string, mood: Mood, audioUrl: string) {
    this.audioUrl = audioUrl;
    this.mood = mood;
    this.bgUrl = bgUrl;
  }
}
