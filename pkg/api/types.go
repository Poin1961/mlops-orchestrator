package api

type Workflow struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       WorkflowSpec `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type WorkflowSpec struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name    string `yaml:"name"`
	Image   string `yaml:"image"`
	Command []string `yaml:"command"`
	Inputs  []Artifact `yaml:"inputs,omitempty"`
	Outputs []Artifact `yaml:"outputs,omitempty"`
}

type Artifact struct {
	Name string `yaml:"name"`
	From string `yaml:"from,omitempty"`
	To   string `yaml:"to,omitempty"`
}
