---
title: "Quick Start: Automation"
sub_link: "quick-start/automation"
meta_title: "Automation - Polyaxon quick start tutorial - Core Concepts"
meta_description: "Automation - Become familiar with the ecosystem of Polyaxon tools with a top-level overview and useful links to get you started."
visibility: public
status: published
tags:
    - tutorials
    - concepts
    - quick-start
sidebar: "intro"
---

> **Note**: DAGs in eager mode are not available yet, this section of the tutorial can only run on Polyaxon EE and Polyaxon Cloud.

Polyaxon provides several features for scaling and automating your process.

There are different ways to scale your operations:
 * Running distributed jobs
 * Running hyperparameter tuning and parallel jobs
 * Automating a large and a complex process

Polyaxon has APIs and clients that you can use with your favorite scheduler.
It also comes with built-in support for distributed jobs, parallel executions, an optimization and a flow engine.

## Automation with DAGs

DAGs are one of the runtimes supported by Polyaxon.

The file `dag.yml` contains a DAG definition, it automates the journey we built manually in this tutorial:

```yaml
version: 1.1
kind: component
name: automated-process
description: runs an experiment, if the loss is higher than max_loss start a hyperparameter tuning process, and then print the best models
inputs:
- {name: max_loss, type: float, value: 0.01, isOptional: true, description: "max loss to start a tuning job."}
- {name: top, type: int, value: 5, isOptional: true, description: "top jobs."}
run:
  kind: dag
  operations:
  - name: build
    params:
      destination:
        connection: docker-connection
        value: polyaxon/polyaxon-quick-start
    runPatch:
      init:
      - dockerfile:
          image: "tensorflow/tensorflow:2.0.1-py3"
          run:
          - 'pip3 install --no-cache-dir -U polyaxon["polyboard","polytune"]'
          langEnv: 'en_US.UTF-8'
  - name: experiment
      urlRef: https://raw.githubusercontent.com/polyaxon/polyaxon-quick-start/master/experimentation/typed.yml
      dependencies: [build]
      params:
        learning_rate : 0.005
        epochs: 10
  - name: tune
    urlRef: https://raw.githubusercontent.com/polyaxon/polyaxon-quick-start/master/experimentation/typed.yml
    params:
      upstream_loss:
        ref: ops.experiment
        value: outputs.loss
        contextOnly: true
    condition: "{{ upstream_loss > dag.inputs.max_loss }}"
    matrix:
      kind: random
      concurrency: 2
      numRuns: 20
      params:
        learning_rate:
          kind: linspace
          value: 0.001:0.1:5
        dropout:
          kind: choice
          value: [0.25, 0.3]
        conv_activation:
          kind: pchoice
          value: [[relu, 0.1], [sigmoid, 0.8]]
        epochs:
          kind: choice
          value: [5, 10]
  - name: best_model
    dependencies: [experiment, tune]
    trigger: all_done
    component:
      run:
        kind: job
        init:
        - git: {url: "https://github.com/polyaxon/polyaxon-quick-start"}
        container:
          image: polyaxon/polyaxon-quick-start
          command: [python3, "{{ globals.artifacts_path }} + /polyaxon-quick-start/best_models.py"]
          args: ["--project={{ globals.project_name }}", "--top={{ dag.inputs.top }}"]
```

This DAG will start an experiment, if the experiment has a `loss > max_loss`
it will start a tuning job based on a random search algorithm,
finally, it will run a container to print the best 5 models.

The DAG itself is parameterized, we can pass different values for `max_loss` and `top`.

```bash
$ polyaxon run --url https://raw.githubusercontent.com/polyaxon/polyaxon-quick-start/master/automation/dag.yml -P loss=0.002 -P top=10
```

A DAG definition is also managed internally by a pipeline, which means you can also leverage all
the pipeline tools to [control the caching](/docs/automation/helpers/cache/) for runs with similar details,
 [concurrency](/docs/automation/helpers/concurrency/) for managing the number of parallel jobs, and [early stopping strategies](/docs/automation/helpers/early-stopping/).

To learn more about [DAGs](/docs/automation/flow-engine/).

## Learn More

You can check the [automation section](/docs/automation/) for more details about all the automation features.
