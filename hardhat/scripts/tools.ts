export const JobSpec = ({
  image = 'ubuntu',
  entrypoint = [],
}: {
  image?: string,
  entrypoint?: string[],
}) => {
  return JSON.stringify({
    Engine: 'docker',
    Verifier: 'noop',
    Publisher: 'estuary',
    Docker: {
      Image: image,
      Entrypoint: entrypoint,
    },
    Outputs: [{
      Name: 'outputs',
      Path: '/outputs',
    }],
    Deal: {
      Concurrency: 1,
    },
  })
}