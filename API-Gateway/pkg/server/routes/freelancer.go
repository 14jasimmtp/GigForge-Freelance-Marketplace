package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Freelancer(
	api fiber.Router,
	profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler,
){
	profiles:=api.Group("/profile")
	profiles.Get("",profile.GetProfile)
	profiles.Post("/education",profile.AddEducationDetails)
	profiles.Patch("/education",profile.UpdateEducation)
	profiles.Post("/experience",profile.AddExperience)
	profiles.Patch("/experience",profile.UpdateExperience)
	profiles.Delete("/experience",profile.RemoveExperience)
	profiles.Delete("/education",profile.DeleteEducation)
	profiles.Post("/description",profile.AddProfileDescription)
	profiles.Patch("/description",profile.EditProfileDescription)
	profiles.Put("/photo",profile.UpdateProfilePhoto)
	profiles.Post("/skill",profile.UpdateSkilltoProfile)

	// jobs:=api.Group("/job")
	// jobs.Get("/:id",job.GetJob)
	// // jobs.Get("",job.GetJobs)
	// jobs.Post("/proposal/:id",job.SendProposal)
	// jobs.Post("/",job.AcceptOffer)
	// jobs.Get("/proposals",job.GetMyProposals)

}

