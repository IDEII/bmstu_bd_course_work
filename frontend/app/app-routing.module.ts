import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegistrationComponent } from './registration/registration.component';
import { LoginComponent } from './login/login.component';
import { CreateTournamentComponent } from './create-tournament/create-tournament.component';
import { AddClubComponent } from './add-club/add-club.component';
import { FighterComponent } from './fighter/fighter.component';
import { TournamentComponent } from './tournament/tournament.component';
import { ClubsComponent } from './clubs/clubs.component';
import { HomeComponent } from './home/home.component';
import { ClubProfileComponent } from './club-profile/club-profile.component';
import { OrganazerProfileComponent } from './organazer-profile/organazer-profile.component';
import { OrganazerListComponent } from './organazer-list/organazer-list.component';
import { TournamentListComponent } from './tournament-list/tournament-list.component';
import { FightersComponent } from './fighters/fighters.component';
import { AddOrganazerComponent } from './add-organazer/add-organazer.component';
import { AddFighterComponent } from './add-fighter/add-fighter.component';
import { RequestManagerComponent } from './request-manager/request-manager.component';
import { UserProfileComponent } from './user-profile/user-profile.component';

const routes: Routes = [
  { path: 'register', component: RegistrationComponent },
  { path: 'login', component: LoginComponent },
  { path: 'add_club', component: AddClubComponent },
  { path: 'add_fighter', component: AddFighterComponent },
  { path: 'add_organazer', component: AddOrganazerComponent },
  { path: 'create_tournament', component: CreateTournamentComponent },
  { path: 'draw', component: LoginComponent },
  { path: 'fighter/:id', component: FighterComponent },
  { path: 'fighters', component: FightersComponent },
  { path: 'tournament/:id', component: TournamentComponent },
  { path: 'home', component: HomeComponent },
  { path: 'clubs', component: ClubsComponent },
  { path: 'tournaments', component: TournamentListComponent },
  { path: 'club-profile/:id', component: ClubProfileComponent },
  { path: 'user-profile/:id', component: UserProfileComponent },
  { path: 'organazer-profile/:id', component: OrganazerProfileComponent },
  { path: 'organazers', component: OrganazerListComponent },
  { path: 'request_manager', component: RequestManagerComponent },
  { path: '', redirectTo: 'home', pathMatch: 'full' }];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
