import Link from 'next/link'

export default function NotFound() {
    return (
        <div>
            <h1>Movie not found</h1>
            <p>We could not find the movie you are looking for.</p>
            <Link href="/movies">Back to movies</Link>
        </div>
    )
}