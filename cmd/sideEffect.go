package cmd

// SideEffect this defines some side effect experienced by the person tied to this side effect
type SideEffect struct {
	IsPositive  bool
	Name        string
	Description string
	Group       string
}

// Person currently is only a logical construct used to hold an array of side effects
type Person struct {
	SideEffects []SideEffect
}

// Notes on Side Effects:
// There are many different types of side effects.
// To simplify things there are really only two types of side effects that we care about.
// 1. Desireable side effects -- Which are those side effects that improve our situation.
// 2. Undesireable side effects -- Which are those side effects that make our situation worse.

// So ideally I would like to catagorize all side effects into those two categories.
// However how do I get people to give me their side effects? Well convincing them is a concern it is not as big a concern as normalizing the data.
// Thus to normalize 'side effects' we will need to catagorize them.
// Which means we will need to start to record all possible side effects...We need a huge list of all Desireable and another of Undesirable side effects.
// This list will grow quickly meaning you will need to segregate it by some method...but we can decide on that later...my ideas right now are centralized around removing a group for Pyscological

// While looking for a list of all side effects I have found the ICD-10 Which is the unified billing codes for the whole world. I will use this to define a standard set of Side Effects.
// I have copied the ICD-10 codes into this repo.

// I am not sure that the ICD-10 codes are going to work the way I want them to work...Maybe we can map them to the side effects that I map...I was hoping to have something more
// along the lines of something like the dsm 5 because that actually tells you how to diagnose not just giving you codes so that you can map treatment to codes so you can charge the patient.
