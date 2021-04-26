# Introduction to AI and AI agent

[TOC]

## General AI knowledge

Why study AI?

1. science
2. IoT
3. Labour
4. Search engines

What is AI?

1. System that think like humans
2. System that think rationally (理性地)
3. System that act like humans
4. System that act rationally

### Turing Test

一台機器是否有智慧？可以讓一個人來問問題判斷，若題問者無法分辨受試者機器及受試者人類，則表示該機器有智慧

#### Full/Total Turing Test

What would a computer need to pass the Full/Total Turing test?

The "Total Turing test" is a variation of the Turing test, proposed by cognitive scientist [Stevan Harnad](https://en.wikipedia.org/wiki/Stevan_Harnad), adds two further requirements to the traditional Turing test. 

The interrogator can also 

1. Test the perceptual abilities (感知能力) of the subject 
   (requiring [computer vision](https://en.wikipedia.org/wiki/Computer_vision)) 

2. The subject's ability to manipulate objects
   (requiring [robotics](https://en.wikipedia.org/wiki/Robotics)).

## Intelligent Agents

An agent is anything that can be viewed as perceiving its environment through sensors and acting upon that environment through effectors. A human agent has eyes, ears, and other organs for sensors, and hands, legs, mouth, and other body parts for effectors. A robotic agent substitutes cameras and infrared range finders for the sensors and various motors for the effectors. A software agent has encoded bit strings as its percepts and actions

**rational agent**

A rational agent is one that does the right thing. 台大于天立：比如一個機械手臂能夾一顆雞蛋不要破就算一個 rational agent

![Agent](./src/agent.png)

### How is an Agent Different from Other Software?

+ Agents are autonomous, i,e,. they act on behalf of the user

+ Agents contain some level of intelligence from fixed rules to learning engines that allow them to adapt to changes in the environment

+ Agents don't only act reactively, but sometimes also proactively
+ Agents have social ability, i.e., they communicate with the user, the system, and other agents as required.
+ Agents may also cooperate with other agents to carry our more complex tasks than they themselves can handle
+ Agents may migrate from one system to another to access remote resources or even to meet other agents

### Structure of Intelligent Agents

Agent = architecture + program

Agent program: the implementation of $f: P^*→A$ 

+ P: perception
+ A: Actions

### Agent Types

+ Simple reflex agents

  Knowing about the current state of the environment

  + 什麼樣的條件就觸發什麼樣的事 Reactive: no memory
  + Benefits: robustness, fast response time
  + Challenges: scalability, how intelligent? and how do you debug them?

+ Model-based reflex agents

  It woks by finding a rule whose condition matches the current situation. A model-based agent can handle partially observable environments by use of model about the world. The agent has o keep track of internal state which is adjusted by each percept and that depends on the percept history.

+ Goal-based agents

  As well as a current state description, the agent needs some sort of goal information, which describes situations that are desirable — for example, being at the passenger's destination. Their every action is intended to reduce its distance from the goal. This allows the agent a way to choose among multiple possibilities, selecting the one which reaches a goal state.

+ Utility-based agents

  基於效益的代理人

  The agents which are developed having their end uses as building blocks are called utility based agents. When there are multiple possible alternatives, then to decide which one is best, utility-based agents are used. They choose actions based on a preference for each state. 

+ Learning agents

  A learning agent in AI is the type of agent which can learn from its past experiences or it has learning capabilities. It starts to act with basic knowledge and then able to act and adapt automatically through learning.

**reference **

+ [Agents in Artificial Intelligence - Geeks](https://www.geeksforgeeks.org/agents-artificial-intelligence/)

## Omniscience

An omniscience (全知) agent knows the actual outcome of its actions and can act accordingly; but omniscience impossible in reality. 

Our definition of a rationality does not require omniscience, then, because the rational choice depends only on the percept sequence to date.

## Properties of Task Environments

### Fully observable v.s. partially observable

If an agent's sensors give it access to the complete state of the environment at each point in time, then we say that the task environment is fully observable.

An environment might be partially observable because of noisy and inaccurate sensors or because parts of the state are simply missing from the sensor data.

If the agent has no sensors at all then the environment is unobservable.

### Deterministic v.s. stochastic

If the next state of the environment is completely determined by the current state and the action executed by the agent, then we say the environment is deterministic; otherwise, it is stochastic

If the environment is partially observable, however, then it could appear to be stochastic.

### Episodic v.s. sequential

In an episodic task environment, the agent's experience is divided into atomic episodes. The next episode does not depend on the actions taken in previous episodes. Many classification tasks are episodic.

### Static v.s. dynamic

If the environment can change while an agent is deliberation, then we say the environment is dynamic for that agent' otherwise, it is static.