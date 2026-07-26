export class Track {
  id: number;
  name: string;
  mood: string;
  bgUrl: string;
  audioUrl: string;

  constructor(id: number, name: string, bgUrl: string, mood: string, audioUrl: string) {
    this.id = id;
    this.name = name;
    this.mood = mood;
    this.audioUrl = audioUrl;
    this.bgUrl = bgUrl;
  }
}
