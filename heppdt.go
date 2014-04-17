package heppdt

// Name returns the name of the default particle data table
func Name() string {
	return defaultTable.Name()
}

// Len returns the size of the default particle data table
func Len() int {
	return defaultTable.Len()
}

// Pdt returns the default particle data table content
func Pdt() map[PID]*Particle {
	return defaultTable.Pdt()
}

// ParticleByID returns the particle information via particle ID
func ParticleByID(pid PID) *Particle {
	return defaultTable.ParticleByID(pid)
}

// ParticleByName returns the particle information via particle name
func ParticleByName(n string) *Particle {
	return defaultTable.ParticleByName(n)
}
