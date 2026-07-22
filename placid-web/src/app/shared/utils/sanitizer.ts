import InvalidInputError from '../exceptions/invalid_input';

export default function sanitizeCredentials(email: string, password: string) {
  if (!email.includes('@') || !email.includes('.') || email.length < 5) {
    throw new InvalidInputError('Please provide a valid email');
  }

  if (password.length < 6) {
    throw new InvalidInputError('Password cannot be less than 6 characters');
  }
}
