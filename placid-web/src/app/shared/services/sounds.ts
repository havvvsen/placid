import { Mood } from '../models/mood';
import SoundScape from '../models/soundscape';

export default async function fetchSoundscapes(): Promise<SoundScape[]> {
  let sounds: SoundScape[] = [
    {
      mood: Mood.Focus,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Anxiety,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Relax,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.TinnutusRelief,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Focus,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Focus,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Focus,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
    {
      mood: Mood.Focus,
      audioUrl: 'http://localhost:8091/audio/stream',
      bgUrl: 'https://imgur/img.png',
    },
  ];

  return sounds;
}
